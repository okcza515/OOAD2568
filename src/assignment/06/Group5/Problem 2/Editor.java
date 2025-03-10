import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.util.ArrayList;
import java.util.List;

public class Editor extends JFrame {
    private DefaultListModel<Note> listModel; // Stores the list of notes
    private JList<Note> notesList; // List UI element
    private Filter filter; // Search field
    private JButton saveButton; // Save button
    private List<Note> allNotes = new ArrayList<>(); // Full list of notes

    // Constructor
    public Editor() {
        createGUI(); // Initialize the UI
    }

    // Method to create the graphical interface
    public void createGUI() {  // Made public to allow access from outside the class
        setTitle("Editor");
        setSize(400, 300);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setLayout(new BorderLayout());

        // Initialize components
        listModel = new DefaultListModel<>();
        notesList = new JList<>(listModel);
        filter = new Filter(this); // Pass this editor instance

        saveButton = new JButton("Save"); // Initialize Save button directly
        saveButton.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                saveChanges(); // Call saveChanges method when button is pressed
            }
        });

        // Add components to frame
        add(filter, BorderLayout.NORTH);
        add(new JScrollPane(notesList), BorderLayout.CENTER);
        add(saveButton, BorderLayout.SOUTH);

        // Load initial data
        loadDummyData();

        setVisible(true);
    }

    // Save changes (dummy function for now)
    public void saveChanges() {
        System.out.println("Changes saved!");
        JOptionPane.showMessageDialog(this, "Changes have been saved successfully.");
    }

    // Load sample data into the list
    private void loadDummyData() {
        allNotes.add(new Note("Meeting Notes", "Discuss project progress"));
        allNotes.add(new Note("Project Plan", "Outline of tasks and deadlines"));
        allNotes.add(new Note("Shopping List", "Buy groceries and supplies"));

        // Add to UI list
        for (Note note : allNotes) {
            listModel.addElement(note);
        }
    }

    // Update the list display when filtering
    public void updateElementsList(DefaultListModel<Note> filteredList) {
        notesList.setModel(filteredList);
    }

    // Getter for allNotes
    public List<Note> getAllNotes() {
        return allNotes;
    }

    // Main method to run the editor
    public static void main(String[] args) {
        // In this case, calling createGUI directly from the constructor of Editor.
        // new Editor(); // This will already invoke createGUI()
        
        // However, if you need to explicitly call createGUI, you can do it as shown below
        Editor editor = new Editor();
        editor.createGUI(); // This will also work, but it's not necessary because the constructor already calls it.
    }
}