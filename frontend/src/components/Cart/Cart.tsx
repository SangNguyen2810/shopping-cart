import { useState, useMemo, useCallback } from 'react';
import { Button, LoadingOverlay } from '@mantine/core';
import { notifications } from '@mantine/notifications';
import useCartStore from '@/store/cartStore';
import useVoucherStore from '@/store/voucherStore';
import useOrderConfirmationStore from '@/store/orderConfirmationStore';
import { Product } from '@/types';
import CartItem from './CartItem';
import styles from './Cart.module.css';
import { NOTIFICATIONS } from '@constants/index';
import { placeOrder } from '@/api/order-api';
import { validateDiscountCode } from '@/api/discount-api';

interface CartProps {
  products: Product[];
}

const Cart = ({ products }: CartProps) => {
  const items = useCartStore((state) => state.items);
  const updateQuantity = useCartStore((state) => state.updateQuantity);
  const clearCart = useCartStore((state) => state.clearCart);
  const setConfirmation = useOrderConfirmationStore((state) => state.setConfirmation);

  const {
    voucherCode,
    voucherApplied,
    discount,
    discountRate,
    voucherError,
    setVoucherCode,
    setVoucherApplied,
    setDiscount,
    setDiscountRate,
    setVoucherError,
    resetVoucher
  } = useVoucherStore();

  const [isLoading, setIsLoading] = useState(false);
  const [isApplyingVoucher, setIsApplyingVoucher] = useState(false);

  const getProduct = useCallback((id: string) => products.find((p) => p.id === id), [products]);

  const cartItems = useMemo(() => {
    return items.map((item) => {
      const product = getProduct(item.productId);
      if (!product) return null;

      return (
        <CartItem
          key={item.productId}
          product={product}
          quantity={item.quantity}
          onRemove={() => updateQuantity(item.productId, 0)}
        />
      );
    });
  }, [items, updateQuantity, getProduct]);

  const totalItems = items.reduce((acc, item) => acc + item.quantity, 0);

  const total = items.reduce((acc, item) => {
    const product = getProduct(item.productId);
    return product ? acc + product.price * item.quantity : acc;
  }, 0);

  const handleApplyVoucher = async () => {
    if (!voucherCode) {
      setVoucherError('Please enter a voucher code');
      notifications.show({
        title: NOTIFICATIONS.ERROR.TITLE,
        message: 'Please enter a voucher code',
        color: NOTIFICATIONS.ERROR.COLOR,
      });
      return;
    }

    setIsApplyingVoucher(true);
    setVoucherError('');
    
    try {
      const response = await validateDiscountCode(voucherCode);
      
      if (response.valid && response.discountRate) {
        setDiscountRate(response.discountRate);
        setDiscount(total * response.discountRate);
        setVoucherApplied(true);
        setVoucherError('');
        notifications.show({
          title: NOTIFICATIONS.SUCCESS.TITLE,
          message: `Voucher applied successfully! You received ${response.discountRate * 100}% discount.`,
          color: NOTIFICATIONS.SUCCESS.COLOR,
        });
      } else {
        const errorMessage = response.error || 'Invalid voucher code';
        setVoucherError(errorMessage);
        notifications.show({
          title: NOTIFICATIONS.ERROR.TITLE,
          message: errorMessage,
          color: NOTIFICATIONS.ERROR.COLOR,
        });
      }
    } catch (error) {
      console.error('Voucher application failed:', error);
      setVoucherError('Failed to apply voucher. Please try again.');
      notifications.show({
        title: NOTIFICATIONS.ERROR.TITLE,
        message: 'Failed to apply voucher. Please try again.',
        color: NOTIFICATIONS.ERROR.COLOR,
      });
    } finally {
      setIsApplyingVoucher(false);
    }
  };

  const finalTotal = total - discount;

  const confirmOrder = async () => {
    setIsLoading(true);
    setVoucherError('');
    
    try {
      const { orderId, items: orderItems } = await placeOrder(
        items, 
        voucherApplied ? voucherCode : undefined
      );

      const orderTotal = orderItems.reduce(
        (sum, item) => sum + item.product.price * item.quantity,
        0,
      );

      const finalAmount = voucherApplied ? orderTotal * (1 - discountRate) : orderTotal;

      setConfirmation({
        orderId,
        totalItems: orderItems.reduce((sum, item) => sum + item.quantity, 0),
        totalAmount: finalAmount,
        items: orderItems,
        voucherCode: voucherApplied ? voucherCode : '',
      });

      resetVoucher();
      clearCart();

      notifications.show({
        title: NOTIFICATIONS.SUCCESS.TITLE,
        message: 'Your order has been placed successfully!',
        color: NOTIFICATIONS.SUCCESS.COLOR,
      });
    } catch (error) {
      console.error('Order confirmation failed:', error);
      
      setVoucherError('');
      
      notifications.show({
        title: NOTIFICATIONS.ERROR.TITLE,
        message: 'Failed to place order. Please try again.',
        color: NOTIFICATIONS.ERROR.COLOR,
      });
    } finally {
      setIsLoading(false);
    }
  };

  const handleVoucherCodeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setVoucherCode(e.target.value);
    if (voucherError) {
      setVoucherError('');
    }
  };

  const handleVoucherButton = () => {
    if (voucherApplied) {
      resetVoucher();
    } else {
      handleApplyVoucher();
    }
  };

  if (items.length === 0) {
    return (
      <article className={styles.cartTotal}>
        <section className="mb-[41px] grid grid-cols-1">
          <header className="mt-[26px] mx-[23px]">
            <h3 className="text-2xl " style={{ color: 'var(--color-text-highlight)' }}>
              Your Cart (0)
            </h3>
          </header>
          <section className="grid grid-cols-1 justify-items-center mt-14 gap-y-7">
            <img src="/total-cart.png" alt="Total Cart" className="w-[100px] h-[100px]" />
            <p>Your added items will appear here</p>
          </section>
        </section>
      </article>
    );
  }

  return (
    <article className={styles.cartTotal}>
      <LoadingOverlay visible={isLoading} />
      <section className="px-6 pt-6">
        <h3
          className="text-[22px] font-medium mb-6"
          style={{ color: 'var(--color-text-highlight)' }}
        >
          Your Cart ({totalItems})
        </h3>

        <ul className={`${styles.cartScroll} space-y-6`}>{cartItems}</ul>
      </section>

      <section className="px-6 pb-6">
        <header
          className="flex justify-between items-center py-4 border-t mt-6"
          style={{ borderColor: 'var(--color-border-light)' }}
        >
          <p className="text-[15px] font-medium">Order Total</p>
          <p className="text-[22px] font-bold">${total.toFixed(2)}</p>
        </header>

        <section className="flex flex-col mb-4">
          <div className="flex items-center gap-2">
            <input
              type="text"
              placeholder="Voucher code"
              value={voucherCode}
              onChange={handleVoucherCodeChange}
              className={`flex-1 px-4 py-2 text-sm border rounded-lg focus:outline-none focus:ring-1 ${
                voucherError ? 'border-red-500 focus:ring-red-500' : 'focus:ring-[#C63B0F]'
              }`}
              style={{ 
                borderColor: voucherError ? '#ef4444' : 'var(--color-border-light)' 
              }}
              disabled={isApplyingVoucher || voucherApplied}
            />
            <Button
              size="xs"
              variant={voucherApplied ? "filled" : "outline"}
              radius="md"
              onClick={handleVoucherButton}
              className={`h-[38px] ${
                voucherApplied 
                  ? 'text-white bg-[#4CAF50] hover:bg-[#3d8b40]' 
                  : 'text-[#C63B0F] border-[#C63B0F] hover:bg-[#FFF5F1]'
              }`}
              loading={isApplyingVoucher}
            >
              {voucherApplied ? 'Reset' : 'Apply'}
            </Button>
          </div>
          {voucherError && (
            <p className="text-red-500 text-xs mt-1">{voucherError}</p>
          )}
        </section>

        {voucherApplied && (
          <section className="flex justify-between items-center mb-4 px-1">
            <p className="text-[13px] text-[#4CAF50] font-medium">Discount Applied</p>
            <p className="text-[15px] text-[#4CAF50] font-medium">-${discount.toFixed(2)} ({discountRate * 100}%)</p>
          </section>
        )}

        {voucherApplied && (
          <header
            className="flex justify-between items-center py-2 mb-4"
            style={{ borderColor: 'var(--color-border-light)' }}
          >
            <p className="text-[15px] font-medium">Final Total</p>
            <p className="text-[22px] font-bold">${finalTotal.toFixed(2)}</p>
          </header>
        )}

        <section
          className="py-[14px] px-4 rounded-lg flex w-full justify-center gap-2 mb-4"
          style={{ backgroundColor: 'var(--color-bg-accent)' }}
        >
          <img src="/carbon-neutral.png" alt="Eco friendly" />
          <div className="flex text-[13px]" style={{ color: '#968D8B' }}>
            <p>This is a &nbsp;</p>
            <p className="font-bold" style={{ color: '#837975' }}>
              carbon-neutral &nbsp;
            </p>
            <p> delivery</p>
          </div>
        </section>

        <Button
          fullWidth
          color="rose"
          radius="xl"
          size="xs"
          onClick={confirmOrder}
          className="transition-all duration-200 ease-in-out hover:-translate-y-0.5 hover:shadow-md h-[48px] !text-[#EFBAA2]"
        >
          Confirm Order
        </Button>
      </section>
    </article>
  );
};

export default Cart;
