package apptheme

import "image/color"

var (
	// Primary color (dark blue) for interface elements
	Primary = color.RGBA{34, 87, 122, 255}

	// Secondary color (green accent) for interactive elements
	Secondary = color.RGBA{46, 204, 113, 255}

	// Tertiary color (light gray) for background elements
	Tertiary = color.RGBA{236, 240, 241, 255}

	// Main text color (dark gray) for primary text
	TextPrimary = color.RGBA{44, 62, 80, 255}

	// Secondary text color (medium gray) for secondary text
	TextSecondary = color.RGBA{127, 140, 141, 255}

	// Warning color (orange) for alert messages
	Warning = color.RGBA{243, 156, 18, 255}

	// Error color (red) for error messages
	Error = color.RGBA{231, 76, 60, 255}

	// Background color (light gray) for the application background
	Background = color.RGBA{250, 250, 250, 255}

	// Color for expense amount (red)
	Expense = color.RGBA{192, 57, 43, 255}

	// Color for income amount (green)
	Income = color.RGBA{39, 174, 96, 255}

	// Color for unclassified amounts (medium gray)
	Unclassified = color.RGBA{127, 140, 141, 255}

	// Color for money transfer between accounts (purple)
	Transfer = color.RGBA{155, 89, 182, 255}

	// Color for savings amount (blue)
	Saving = color.RGBA{52, 152, 219, 255}
)
