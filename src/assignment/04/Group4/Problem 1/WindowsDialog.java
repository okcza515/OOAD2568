public class WindowsDialog implements DialogFactory{
	@Override
	public void renderWindow() {
		Button okButton = createButton();
		okButton.render();
	}

	@Override
	public Button createButton() {
		return new WindowsButton();
	}
}

// Kanitsorn Darunaitorn 65070501069