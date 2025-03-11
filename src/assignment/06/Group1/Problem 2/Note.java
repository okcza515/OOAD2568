public class Note {
	private String name;
	private String text;

	public Note() {
		name = "New note";
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getText() {
		return text;
	}

	public void setText(String text) {
		this.text = text;
	}

	public String toString() {
		return name;
	}
}
