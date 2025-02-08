public abstract class Dialog {

    protected void renderWindow(){
        Button okButton = createButton();
        okButton.render();
    };

    protected abstract Button createButton();

}

// 65070501085