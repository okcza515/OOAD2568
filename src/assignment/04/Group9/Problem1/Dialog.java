//Chayaphon Chaisangkha 65070503409
// Dialog.java
public abstract class Dialog {
    public void renderWindow() {
        // Template method
        Button button = createButton();
        button.render();
    }
    
    // Factory method
    public abstract Button createButton();
}