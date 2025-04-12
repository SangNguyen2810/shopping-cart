import { Modal, ScrollArea, Image, Button } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import useOrderConfirmationStore from '@/store/orderConfirmationStore';
import useCartStore from '@/store/cartStore';
import { OrderResponseItem } from '@/api/order-api';
import { useEffect, useRef } from 'react';

const OrderConfirmation = () => {
  const [opened, { open, close }] = useDisclosure(false);
  const { confirmation, setConfirmation } = useOrderConfirmationStore();
  const clearCart = useCartStore((state) => state.clearCart);
  const modalRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (confirmation) {
      console.log('Opening modal with confirmation:', confirmation);
      open();
    }
  }, [confirmation, open]);

  const handleClose = () => {
    console.log('Closing modal');
    close();
    setConfirmation(null);
    clearCart();
  };

  useEffect(() => {
    if (modalRef.current && opened) {
      setTimeout(() => {
        const button = modalRef.current?.querySelector('button');
        button?.scrollIntoView({ behavior: 'smooth', block: 'center' });
      }, 400);
    }
  }, [opened]);

  return (
    <Modal
      id="order-confirmation-modal"
      opened={!!confirmation && opened}
      onClose={handleClose}
      centered
      size="lg"
      withCloseButton={false}
      overlayProps={{
        color: '#000',
        opacity: 0.75,
        blur: 5,
      }}
      transitionProps={{
        transition: 'fade',
        duration: 300,
      }}
      radius="md"
      closeOnClickOutside={true}
      closeOnEscape={true}
      scrollAreaComponent={ScrollArea}
    >
      {confirmation && (
        <article className="mt-10 mx-2" ref={modalRef}>
          <button>
            <Image src="/confirmed.png" alt="Order Confirmed" />
          </button>

          <div className="mt-7">
            <header className="pb-2 text-[40px] font-semibold">Order Confirmed</header>

            <h2 className="pb-[30px] text-[#A49B99] text-base">We hope you enjoy your food!</h2>

            <article className="bg-gray-50 rounded-lg">
              {confirmation.items && confirmation.items.length > 0 ? (
                <>
                  {confirmation.items.map((item: OrderResponseItem, index: number) => (
                    <section key={index} className="grid grid-cols-4 gap-x-4 pt-7 pb-4 mx-7 border-b border-gray-200">
                      <Image
                        src={item.product.image?.thumbnail}
                        alt={item.product.name}
                        className="w-[48px] h-auto col-span-1"
                      />
                      <div className="col-span-2 grid">
                        <h2 className="">{item.product.name}</h2>
                        <span className="text-sm flex items-center gap-x-4">
                          <p style={{ color: 'var(--color-text-highlight-dimmed)' }}>
                            {item.quantity}x
                          </p>
                          <div className="flex items-center gap-x-1">
                            <p style={{ color: 'var(--color-text-highlight-dimmed)' }}>@</p>
                            <p style={{ color: 'var(--color-text-secondary)' }}>
                              ${item.product.price.toFixed(2)}
                            </p>
                          </div>
                        </span>
                      </div>
                      <div className="flex items-center justify-end">
                        <p className="font-medium">
                          ${(item.quantity * item.product.price).toFixed(2)}
                        </p>
                      </div>
                    </section>
                  ))}
                </>
              ) : (
                <>
                  <p>Order ID: {confirmation.orderId}</p>
                  <p>Total Items: {confirmation.totalItems}</p>
                  <p className="mt-2 font-medium">
                    Total Amount: ${confirmation.totalAmount.toFixed(2)}
                  </p>
                </>
              )}

              <footer className="flex justify-between mx-7 py-6">
                <p className="font-medium">Order Total</p>
                <p className="font-bold text-lg">${confirmation.totalAmount.toFixed(2)}</p>
              </footer>
            </article>

            <Button
              fullWidth
              color="red"
              radius="xl"
              size="xs"
              onClick={handleClose}
              className="transition-all  !text-[#EFBAA2] duration-200 ease-in-out hover:-translate-y-0.5 hover:shadow-md bg-[#C73B0E] hover:bg-[#B94825] h-[48px] mt-5"
            >
              Start New Order
            </Button>
          </div>
        </article>
      )}
    </Modal>
  );
};

export default OrderConfirmation;
