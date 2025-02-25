//65070501019 Natlada Simasathien
public class WindowsDialog implements DialogFactory {
    @Override
    public void renderWindow() {
        WindowsButton okButton = createButton();
        okButton.render();
    }
    @Override
	public WindowsButton createButton() {
		return new WindowsButton();
	}
}