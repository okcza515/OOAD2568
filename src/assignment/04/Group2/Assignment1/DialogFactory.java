public abstract class DialogFactory {
    public Dialog CreateDialog() {
        return BuildDialog();
    }

    protected abstract Dialog BuildDialog();
}
