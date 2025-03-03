import java.io.File;

public class Application {
    public static void main(String[] args){
        VideoConversionFacade converter = new VideoConversionFacade();
        String filename = "example.ogg";
        String outputFormat = "mp4";

        File result = converter.convertVideo(filename,outputFormat);

        System.out.println("Conversion complete. File: "+ result.getName());
    }
}
