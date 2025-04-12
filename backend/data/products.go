package data

import "shopping-cart/backend/models"

var Products = []models.Product{
	{
		ID:   "1",
		Name: "Waffle with Berries",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
		},
		Category: "Waffle",
		Price:    6.5,
	},
	{
		ID:   "2",
		Name: "Vanilla Bean Crème Brûlée",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-creme-brulee-desktop.jpg",
		},
		Category: "Crème Brûlée",
		Price:    7,
	},
	{
		ID:   "3",
		Name: "Macaron Mix of Five",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-macaron-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-macaron-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-macaron-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-macaron-desktop.jpg",
		},
		Category: "Macaron",
		Price:    8,
	},
	{
		ID:   "4",
		Name: "Classic Tiramisu",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-tiramisu-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-tiramisu-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-tiramisu-desktop.jpg",
		},
		Category: "Tiramisu",
		Price:    5.5,
	},
	{
		ID:   "5",
		Name: "Pistachio Baklava",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-baklava-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-baklava-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-baklava-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-baklava-desktop.jpg",
		},
		Category: "Baklava",
		Price:    4,
	},
	{
		ID:   "6",
		Name: "Lemon Meringue Pie",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-meringue-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-meringue-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-meringue-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-meringue-desktop.jpg",
		},
		Category: "Pie",
		Price:    5,
	},
	{
		ID:   "7",
		Name: "Red Velvet Cake",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-cake-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-cake-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-cake-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-cake-desktop.jpg",
		},
		Category: "Cake",
		Price:    4.5,
	},
	{
		ID:   "8",
		Name: "Salted Caramel Brownie",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-brownie-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-brownie-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-brownie-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-brownie-desktop.jpg",
		},
		Category: "Brownie",
		Price:    4.5,
	},
	{
		ID:   "9",
		Name: "Vanilla Panna Cotta",
		Image: models.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-panna-cotta-desktop.jpg",
		},
		Category: "Panna Cotta",
		Price:    6.5,
	},
}
