
public class WindowsDialog {
	public void renderWindow() {
		WindowsButton okButton = createButton();
		okButton.render();
	}

	public WindowsButton createButton() {
		return new WindowsButton();
	}
}
