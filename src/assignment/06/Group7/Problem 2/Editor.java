
import javax.swing.*;
import javax.swing.border.LineBorder;
import java.awt.*;

public class Editor {
	
	private static Title title;
    private static TextBox textBox;
    private static AddButton add;
    private static DeleteButton del;
    private static SaveButton save;
    private static List list;
	private static Filter filter;
	private static JLabel titleLabel = new JLabel("Title:");
    private static JLabel textLabel = new JLabel("Text:");
    private static JLabel label = new JLabel("Add or select existing note to proceed...");
	
	public Editor(){
		title = new Title();
		textBox = new TextBox();
		add = new AddButton();
		del = new DeleteButton();
		save = new SaveButton();
		list = new List(new DefaultListModel());
		this.list.addListSelectionListener(listSelectionEvent -> {
            Note note = (Note)list.getSelectedValue();
            if (note != null) {
                getInfoFromList(note);
            } else {
                clear();
            }
        });
		filter = new Filter();
	}
	
	public static void getInfoFromList(Note note) {
        title.setText(note.getName().replace('*', ' '));
        textBox.setText(note.getText());
    }
	
	public static void clear() {
        title.setText("");
        textBox.setText("");
    }
	
	public static void hideElements(boolean flag) {
        titleLabel.setVisible(!flag);
        textLabel.setVisible(!flag);
        title.setVisible(!flag);
        textBox.setVisible(!flag);
        save.setVisible(!flag);
        label.setVisible(flag);
    }
	
	public static void addNewNote(Note note) {
		title.setText("");
        textBox.setText("");
        list.addElement(note);
	}
	
	public static void sendToFilter(ListModel listModel) {
        filter.setList(listModel);
    }
	
	public static void setElementsList(ListModel listM) {
        list.setModel(listM);
        list.repaint();
    }
	
	public static void markNote() {
        try {
            Note note = list.getCurrentElement();
            String name = note.getName();
            if (!name.endsWith("*")) {
                note.setName(note.getName() + "*");
            }
            list.repaint();
        } catch (NullPointerException ignored) {}
    }
	
	public static void deleteNote() {
        list.deleteElement();
    }
	
	public static void saveChanges() {
        try {
            Note note = (Note) list.getSelectedValue();
            note.setName(title.getText());
            note.setText(textBox.getText());
            list.repaint();
        } catch (NullPointerException ignored) {}
    }
	
	public void createGUI() {
		JFrame notes = new JFrame("Notes");
		notes.setSize(960, 600);
        notes.setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
        JPanel left = new JPanel();
        left.setBorder(new LineBorder(Color.BLACK));
        left.setSize(320, 600);
        left.setLayout(new BoxLayout(left, BoxLayout.Y_AXIS));
        JPanel filterPanel = new JPanel();
        filterPanel.add(new JLabel("Filter:"));
        filter.setColumns(20);
        filterPanel.add(filter);
        filterPanel.setPreferredSize(new Dimension(280, 40));
        JPanel listPanel = new JPanel();
        list.setFixedCellWidth(260);
        listPanel.setSize(320, 470);
        JScrollPane scrollPane = new JScrollPane(list);
        scrollPane.setPreferredSize(new Dimension(275, 410));
        listPanel.add(scrollPane);
        JPanel buttonPanel = new JPanel();
        add.setPreferredSize(new Dimension(85, 25));
        buttonPanel.add(add);
        del.setPreferredSize(new Dimension(85, 25));
        buttonPanel.add(del);
        buttonPanel.setLayout(new FlowLayout());
        left.add(filterPanel);
        left.add(listPanel);
        left.add(buttonPanel);
        JPanel right = new JPanel();
        right.setLayout(null);
        right.setSize(640, 600);
        right.setLocation(320, 0);
        right.setBorder(new LineBorder(Color.BLACK));
        titleLabel.setBounds(20, 4, 50, 20);
        title.setBounds(60, 5, 555, 20);
        textLabel.setBounds(20, 4, 50, 130);
        textBox.setBorder(new LineBorder(Color.DARK_GRAY));
        textBox.setBounds(20, 80, 595, 410);
        save.setBounds(270, 535, 80, 25);
        label.setFont(new Font("Verdana", Font.PLAIN, 22));
        label.setBounds(100, 240, 500, 100);
        right.add(label);
        right.add(titleLabel);
        right.add(title);
        right.add(textLabel);
        right.add(textBox);
        right.add(save);
        notes.setLayout(null);
        notes.getContentPane().add(left);
        notes.getContentPane().add(right);
        notes.setResizable(false);
        notes.setLocationRelativeTo(null);
        notes.setVisible(true);
	}
}
