//65070501078 Nawaphon Promnan
public class HtmlDialog implements DialogFactory {
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