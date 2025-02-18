
public class HtmlDialog {

	public void renderWindow() {
		HtmlButton okButton = createButton();
		okButton.render();
	}

	public HtmlButton createButton() {
		return new HtmlButton();
	}
}
