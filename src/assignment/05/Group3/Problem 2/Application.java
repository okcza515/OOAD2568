import java.io.File;

public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();
        File mp4Video = converter.convertVideo("skibidi.ogg", "mp4");

        System.out.println("Convert file completed : " + mp4Video.getAbsolutePath());
    }
}