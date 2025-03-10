public class Note {
    private String name;
    private String text;

    // Constructor
    public Note() {
        this.name = "Untitled";
        this.text = "";
    }

    public Note(String name, String text) {
        this.name = name;
        this.text = text;
    }

    // ✅ Getter for name
    public String getName() {
        return name;
    }

    // ✅ Fix: Add setName(String) method
    public void setName(String name) {
        this.name = name;
    }

    // ✅ Getter for text
    public String getText() {
        return text;
    }

    // ✅ Fix: Add setText(String) method
    public void setText(String text) {
        this.text = text;
    }

    @Override
    public String toString() {
        return name; // Ensures the JList displays the note name
    }
}
