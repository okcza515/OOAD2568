
import javax.swing.*;
import javax.swing.event.DocumentEvent;
import javax.swing.event.DocumentListener;

public class Filter extends JTextField {
    private ListModel<Note> listModel;
    private Editor editor; // Reference to Editor for UI updates

    // Constructor
    public Filter(Editor editor) {
        this.editor = editor;

        // Add a listener to update search results dynamically
        this.getDocument().addDocumentListener(new DocumentListener() {
            @Override
            public void insertUpdate(DocumentEvent e) {
                searchElements(getText());
            }

            @Override
            public void removeUpdate(DocumentEvent e) {
                searchElements(getText());
            }

            @Override
            public void changedUpdate(DocumentEvent e) {
                searchElements(getText());
            }
        });
    }

    // Set the list model (dependency injection)
    public void setList(ListModel<Note> listModel) {
        this.listModel = listModel;
    }

    // Search functionality
    public void searchElements(String query) {
        if (listModel == null) return; // Avoid null pointer exceptions

        DefaultListModel<Note> filteredList = new DefaultListModel<>();

        for (int i = 0; i < listModel.getSize(); i++) {
            Note note = listModel.getElementAt(i);
            if (note.getName().toLowerCase().contains(query.toLowerCase())) {
                filteredList.addElement(note);
            }
        }

        // Update the UI through Editor
        editor.updateElementsList(filteredList);
    }
}
