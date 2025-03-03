//Natchanon 65070501018
public class Main {
    public static void main(String[] args) {
        Application app = new Application();
        File convertedFile = app.convert("example", "mp4");
        System.out.println(convertedFile.getName());
    }
}
