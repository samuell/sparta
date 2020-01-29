package gui

import (
	"sparta/file"
	"sparta/file/encrypt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// TODO: Multi user support by labling the data file exercises-user.xml

// SettingsView contains the gui information for the settings screen.
func SettingsView(window fyne.Window, app fyne.App, XMLData *file.Data, dataLabel *widget.Label) fyne.CanvasObject {

	// TODO: Add setting for changing language.

	// Make it possible for the user to switch themes.
	themeSwitcher := widget.NewSelect([]string{"Dark", "Light"}, func(selected string) {
		switch selected {
		case "Dark":
			app.Settings().SetTheme(theme.DarkTheme())
		case "Light":
			app.Settings().SetTheme(theme.LightTheme())
		}

		// Set the theme to the selected one.
		config.Theme = selected

		// Write the theme update to the file and do it concurrently to not stop the thread.
		go config.Write()
	})

	// Default theme is dark and thus we set the placeholder to that and then refresh it (without a refresh, it doesn't show until hovering on to widget).
	themeSwitcher.PlaceHolder = config.Theme
	themeSwitcher.Refresh()

	// Add the theme switcher next to a label.
	themeSettings := fyne.NewContainerWithLayout(layout.NewGridLayout(2), widget.NewLabel("Change theme"), themeSwitcher)

	// Create the entry for updating the password.
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("New password")

	// Create the button used for actually changing the password.
	passwordButton := widget.NewButtonWithIcon("Change password", theme.ConfirmIcon(), func() {
		// Check that the password is valid.
		if len(passwordEntry.Text) < 8 {
			dialog.ShowInformation("Please enter a valid password", "Passwords need to be at least eight characters long.", window)
			return
		}

		// Ask the user to confirm what we are about to do.
		dialog.ShowConfirm("Are you sure that you want to continue?", "The action will permanently change your password.", func(change bool) {
			if change {
				// Calculate the new PasswordKey.
				PasswordKey = encrypt.EncryptionKey(UserName, passwordEntry.Text)

				// Clear out the text inside the label.
				passwordEntry.SetText("")

				// Write the data encrypted using the new key and do so concurrently.
				go XMLData.Write(&PasswordKey)
			}
		}, window)
	})

	// Create a button for clearing the data of a given profile.
	deleteButton := widget.NewButtonWithIcon("Remove all activities", theme.DeleteIcon(), func() {

		// Ask the user to confirm what we are about to do.
		dialog.ShowConfirm("Are you sure that you want to continue?", "Deleting your data will remove all of your exercises and activities.", func(remove bool) {
			if remove {
				// Run the delete function and do it concurrently to avoid stalling the thread with file io.
				go XMLData.Delete()

				// Clear all the data inside the data label so we don't display the old data.
				dataLabel.SetText("No exercieses have been created yet.")
			}
		}, window)
	})

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), themeSettings, layout.NewSpacer(), fyne.NewContainerWithLayout(layout.NewGridLayout(2), passwordEntry, passwordButton), layout.NewSpacer(), deleteButton)
}