
public class WindowsDialog implements Dialog{
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

// 65070501023 Thanaphol Thangthaweesuk
