
public class HtmlDialog extends Dialog {

	public void renderWindow() {
		HtmlButton okButton = createButton();
		okButton.render();
	}

	public HtmlButton createButton() {
		return new HtmlButton();
	}
}
// Ratchanon Tarawan 65070503464