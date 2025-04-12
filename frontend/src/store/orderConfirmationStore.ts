import { create } from 'zustand';
import { OrderResponseItem } from '@/api/order-api';

interface OrderConfirmation {
  orderId: string;
  totalItems: number;
  totalAmount: number;
  items?: OrderResponseItem[];
  voucherCode?: string;
}

type OrderConfirmationState = {
  confirmation: OrderConfirmation | null;
  setConfirmation: (confirmation: OrderConfirmation | null) => void;
}

export interface OrderConfirmationStore {
  confirmation: OrderConfirmation | null;
  setConfirmation: (confirmation: OrderConfirmation) => void;
  clearConfirmation: () => void;
}

const useOrderConfirmationStore = create<OrderConfirmationState>((set) => ({
  confirmation: null,
  setConfirmation: (confirmation) => {
    console.log('Setting confirmation in store:', confirmation);
    set({ confirmation });
  },
}));

export default useOrderConfirmationStore;