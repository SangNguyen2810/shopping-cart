import { create } from "zustand";

export interface CartItem {
	productId: string;
	quantity: number;
}

export interface CartStore {
	items: CartItem[];
	addItem: (productId: string) => void;
	removeItem: (productId: string) => void;
	updateQuantity: (productId: string, quantity: number) => void;
	clearCart: () => void;
}

const useCartStore = create<CartStore>((set) => ({
	items: [],
	addItem: (productId) =>
		set((state) => {
			const existing = state.items.find((item) => item.productId === productId);
			if (existing) {
				return {
					items: state.items.map((item) =>
						item.productId === productId
							? { ...item, quantity: item.quantity + 1 }
							: item
					),
				};
			}
			return { items: [...state.items, { productId, quantity: 1 }] };
		}),
	removeItem: (productId) =>
		set((state) => {
			const existing = state.items.find((item) => item.productId === productId);
			if (existing) {
				if (existing.quantity > 1) {
					return {
						items: state.items.map((item) =>
							item.productId === productId
								? { ...item, quantity: item.quantity - 1 }
								: item
						),
					};
				} else {
					return {
						items: state.items.filter((item) => item.productId !== productId),
					};
				}
			}
			return state;
		}),
	updateQuantity: (productId, quantity) =>
		set((state) => {
			if (quantity <= 0) {
				return {
					items: state.items.filter((item) => item.productId !== productId),
				};
			}
			return {
				items: state.items.map((item) =>
					item.productId === productId ? { ...item, quantity } : item
				),
			};
		}),
	clearCart: () => set({ items: [] }),
}));

export default useCartStore;
