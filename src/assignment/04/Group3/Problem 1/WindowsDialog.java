
public class WindowsDialog implements Dialog{
	public void renderWindow() {
		WindowsButton okButton = createButton();
		okButton.render();
	}

	public WindowsButton createButton() {
		return new WindowsButton();
	}
}

// 65070501023 Thanaphol Thangthaweesuk
