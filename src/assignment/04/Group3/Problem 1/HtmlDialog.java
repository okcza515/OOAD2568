
public class HtmlDialog implements Dialog {

	public void renderWindow() {
		HtmlButton okButton = createButton();
		okButton.render();
	}

	public HtmlButton createButton() {
		return new HtmlButton();
	}
}

// 65070501023 Thanaphol Thangthaweesuk
