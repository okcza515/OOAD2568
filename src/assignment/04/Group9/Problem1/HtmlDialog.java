// Chanawat Limpanatewin 65070503445
// HtmlDialog.java
public class HtmlDialog extends Dialog {
    @Override
    public Button createButton() {
        return new HtmlButton();
    }
}