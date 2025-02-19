public class WindowsDialog extends Dialog {
	public void renderWindow() {
		WindowsButton okButton = createButton();
		okButton.render();
	}

	public WindowsButton createButton() {
		return new WindowsButton();
	}
}
// Ratchanon Tarawan 65070503464