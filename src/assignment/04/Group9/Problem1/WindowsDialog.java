// WindowsDialog.java //Paratthakon Suksukhon 65070503457
public class WindowsDialog extends Dialog {
    @Override
    public Button createButton() {
        return new WindowsButton();
    }
}