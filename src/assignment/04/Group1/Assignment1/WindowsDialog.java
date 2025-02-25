public class WindowsDialog implements IDialog {
	@Override
	public void renderWindow() {
		IButton okButton = createButton();
		okButton.render();
	}

	@Override
	public IButton createButton() {
		return new WindowsButton();
	}
}
