
public class HtmlDialog implements Dialog {

	@Override
	public void renderWindow() {
		HtmlButton okButton = createButton();
		okButton.render();
	}

	@Override
	public HtmlButton createButton() {
		return new HtmlButton();
	}
}

// 65070501023 Thanaphol Thangthaweesuk
