import java.io.File;

public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();
        File mp4File = converter.convertVideo("video.ogg", "mp4");
        System.out.println("Converted File: " + mp4File.getName());
    }
}