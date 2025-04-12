import { DEFAULT_IMAGE } from "@/constants";

export const getFallbackImage = (category?: string) => {
  if (!category) return DEFAULT_IMAGE;

  switch (category.toLowerCase()) {
    case 'waffle':
      return 'https://placehold.co/400x300?text=Waffle';
    case 'tiramisu':
      return 'https://placehold.co/400x300?text=Tiramisu';
    case 'baklava':
      return 'https://placehold.co/400x300?text=Baklava';
    case 'pie':
      return 'https://placehold.co/400x300?text=Pie';
    case 'cake':
      return 'https://placehold.co/400x300?text=Cake';
    case 'brownie':
      return 'https://placehold.co/400x300?text=Brownie';
    default:
      return DEFAULT_IMAGE;
  }
};