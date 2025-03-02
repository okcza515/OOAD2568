
import java.io.File;

public class Application {
    public static void main(String[] args) {
        VideoConversionFacade converter = new VideoConversionFacade();
        
        // Convert OGG to MP4
        String inputFile = "sample.ogg";
        String outputFormat = "mp4";
        System.out.println("Starting video conversion...");
        File convertedFile = converter.convertVideo(inputFile, outputFormat);

        //Display converted file
        System.out.println("Conversion completed. Output file: " + convertedFile.getName());
    }
}

//Supanut Wongtanom 65070503437
