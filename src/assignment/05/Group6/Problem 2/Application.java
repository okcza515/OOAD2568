import java.io.File;

public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();
        File convertedFile = converter.convertVideo("test.ogg", "mp4");
        System.out.println("Conversion completed: " + convertedFile.getName());
    }
}
