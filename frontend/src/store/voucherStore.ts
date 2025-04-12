import { create } from "zustand";

export interface VoucherStore {
  voucherCode: string;
  voucherApplied: boolean;
  discount: number;
  discountRate: number;
  voucherError: string;
  setVoucherCode: (code: string) => void;
  setVoucherApplied: (applied: boolean) => void;
  setDiscount: (discount: number) => void;
  setDiscountRate: (rate: number) => void;
  setVoucherError: (error: string) => void;
  resetVoucher: () => void;
}

const useVoucherStore = create<VoucherStore>((set) => ({
  voucherCode: '',
  voucherApplied: false,
  discount: 0,
  discountRate: 0,
  voucherError: '',
  setVoucherCode: (code) => set({ voucherCode: code }),
  setVoucherApplied: (applied) => set({ voucherApplied: applied }),
  setDiscount: (discount) => set({ discount }),
  setDiscountRate: (discountRate) => set({ discountRate }),
  setVoucherError: (error) => set({ voucherError: error }),
  resetVoucher: () => set({
    voucherCode: '',
    voucherApplied: false,
    discount: 0,
    discountRate: 0,
    voucherError: ''
  })
}));

export default useVoucherStore; 