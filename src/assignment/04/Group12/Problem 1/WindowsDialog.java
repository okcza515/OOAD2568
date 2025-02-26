
public class WindowsDialog implements IDialog {
	@Override
	public void renderWindow() {
		WindowsButton okButton = createButton();
		okButton.render();
	}

	@Override
	public WindowsButton createButton() {
		return new WindowsButton();
	}
}
