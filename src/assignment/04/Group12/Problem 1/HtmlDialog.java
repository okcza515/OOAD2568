
public class HtmlDialog implements IDialog{
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
