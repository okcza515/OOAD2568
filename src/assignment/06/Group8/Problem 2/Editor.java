import javax.swing.*;
import javax.swing.border.LineBorder;
import java.awt.*;

public class Editor implements EditorMediator {
    private Title title;
    private TextBox textBox;
    private AddButton add;
    private DeleteButton del;
    private SaveButton save;
    private List list;
    private Filter filter;
    private JLabel titleLabel = new JLabel("Title:");
    private JLabel textLabel = new JLabel("Text:");
    private JLabel label = new JLabel("Add or select an existing note to proceed...");

    public Editor() {
        title = new Title(this);
        textBox = new TextBox(this);
        add = new AddButton(this);
        del = new DeleteButton(this);
        save = new SaveButton(this);
        list = new List(new DefaultListModel(), this);
        filter = new Filter(this);
    }

    @Override
    public void addNote() {
        list.addElement(new Note());
    }

    @Override
    public void deleteNote() {
        list.deleteElement();
    }

    @Override
    public void saveNote() {
        try {
            Note note = (Note) list.getSelectedValue();
            note.setName(title.getText());
            note.setText(textBox.getText());
            list.repaint();
        } catch (NullPointerException ignored) {}
    }

    @Override
    public void selectNote(Note note) {
        title.setText(note.getName().replace('*', ' '));
        textBox.setText(note.getText());
    }

    @Override
    public void clearNote() {
        title.setText("");
        textBox.setText("");
    }

    @Override
    public void markNote() {
        try {
            Note note = list.getCurrentElement();
            String name = note.getName();
            if (!name.endsWith("*")) {
                note.setName(note.getName() + "*");
            }
            list.repaint();
        } catch (NullPointerException ignored) {}
    }

    @Override
    public void sendToFilter(ListModel listModel) {
        filter.setList(listModel);
    }

    @Override
    public void setElementsList(ListModel listM) {
        list.setModel(listM);
        list.repaint();
    }

    @Override
    public void hideElements(boolean flag) {
        titleLabel.setVisible(!flag);
        textLabel.setVisible(!flag);
        title.setVisible(!flag);
        textBox.setVisible(!flag);
        save.setVisible(!flag);
        label.setVisible(flag);
    }
    

    public void createGUI() {
        JFrame notes = new JFrame("Notes");
        notes.setSize(960, 600);
        notes.setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
        JPanel left = new JPanel();
        left.setBorder(new LineBorder(Color.BLACK));
        left.setLayout(new BoxLayout(left, BoxLayout.Y_AXIS));
        JPanel filterPanel = new JPanel();
        filterPanel.add(new JLabel("Filter:"));
        filter.setColumns(20);
        filterPanel.add(filter);
        left.add(filterPanel);
        JPanel listPanel = new JPanel();
        JScrollPane scrollPane = new JScrollPane(list);
        scrollPane.setPreferredSize(new Dimension(275, 410));
        listPanel.add(scrollPane);
        left.add(listPanel);
        JPanel buttonPanel = new JPanel();
        buttonPanel.add(add);
        buttonPanel.add(del);
        left.add(buttonPanel);
        JPanel right = new JPanel();
        right.setLayout(null);
        right.setBorder(new LineBorder(Color.BLACK));
        right.add(titleLabel);
        right.add(title);
        right.add(textLabel);
        right.add(textBox);
        right.add(save);
        notes.getContentPane().add(left);
        notes.getContentPane().add(right);
        notes.setResizable(false);
        notes.setLocationRelativeTo(null);
        notes.setVisible(true);
    }
}
