
import javax.swing.*;

public class List extends JList{
	
	private final DefaultListModel LIST_MODEL;

    private EditorMediators mediators;
	
	public List(DefaultListModel listModel, EditorMediators mediators) {
        super(listModel);
        this.LIST_MODEL = listModel;
        this.mediators = mediators;
        setModel(listModel);
        this.setLayoutOrientation(JList.VERTICAL);
        Thread thread = new Thread(new Hide(this));
        thread.start();
    }
	
	public void addElement(Note note) {
        LIST_MODEL.addElement(note);
        int index = LIST_MODEL.size() - 1;
        setSelectedIndex(index);
        ensureIndexIsVisible(index);
        mediators.sendToFilter(LIST_MODEL);
    }
	
	public void deleteElement() {
        int index = this.getSelectedIndex();
        try {
            LIST_MODEL.remove(index);
            mediators.sendToFilter(LIST_MODEL);
        } catch (ArrayIndexOutOfBoundsException ignored) {}
    }
	
	public Note getCurrentElement() {
        return (Note)getSelectedValue();
    }
	
	private class Hide implements Runnable {
        private List list;

        Hide(List list) {
            this.list = list;
        }

        @Override
        public void run() {
            while (true) {
                try {
                    Thread.sleep(300);
                } catch (InterruptedException ex) {
                    ex.printStackTrace();
                }
                if (list.isSelectionEmpty()) {
                    mediators.hideElements(true);
                } else {
                    mediators.hideElements(false);
                }
            }
        }
    }
}

// Sikares Nuntipatsakul 65070503439