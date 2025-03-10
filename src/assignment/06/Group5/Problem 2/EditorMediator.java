class EditorMediator implements Mediator {
    private Title title;
    private TextBox textBox;
    private List list;

    public void registerComponents(Title title, TextBox textBox, List list) {
        this.title = title;
        this.textBox = textBox;
        this.list = list;
    }

    @Override
    public void addNote() {
        list.addElement(new Note());
    }

    @Override
    public void deleteNote() {
        list.removeSelectedElement();
    }

    @Override
    public void saveChanges() {
        Note selected = list.getSelectedElement();
        if (selected != null) {
            selected.setName(title.getText());
            selected.setText(textBox.getText());
        }
    }

    @Override
    public void markNote() {
        list.highlightSelectedElement();
    }
}
