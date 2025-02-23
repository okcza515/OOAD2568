
public class HtmlDialog implements DialogFactory {
    @Override
    public void renderWindow() {
        Button okButton = createButton();
        okButton.render();
    }
    
    @Override
    public Button createButton() {
        return new HtmlButton();
    }
}
