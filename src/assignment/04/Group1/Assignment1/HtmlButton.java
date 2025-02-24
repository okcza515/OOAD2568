public class HtmlButton implements IButton {
    @Override
    public void render() {
    System.out.println("<button>Test Button</button>");
    onClick();
    }
    
    @Override
    public void onClick() {
    System.out.println("Click! Button says - 'Hello World!'");
    }
    }