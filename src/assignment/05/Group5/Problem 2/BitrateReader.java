//Natchanon 65070501018
class BitrateReader {
    public static String read(String filename, String codec){
        System.out.println("...Reading " + codec +  " of \""+ filename +"\"");
        return filename;
    }

    public static String convert(String filename, Codec codec){
        System.out.println("...Converting " + filename +  " to "+ codec.fileType());
        return filename + "." + codec.fileType();
    }    
}
