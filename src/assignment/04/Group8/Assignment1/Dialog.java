abstract class Dialog {
    abstract Button createButton();

    public void renderWindow() {
        Button okButton = createButton();
        okButton.render();
    }
}
