import java.util.Base64;

public class EncryptionDecorator extends DataSourceDecorator {

    private encode(String data) {
        byte[] result = data.getBytes();
        for (int i = 0; i < result.length; i++) {
            result[i] += (byte) 1;
        }
        return Base64.getEncoder().encodeToString(result);
    }

    private decode(String data) {
        byte[] result = Base64.getDecoder().decode(data);
        for (int i = 0; i < result.length; i++) {
            result[i] -= (byte) 1;
        }
        return new String(result);
    }

    public EncryptionDecorator(DataSource source) {
        super(source);
        encryption = new Encryption();
    }

    @Override
    public void writeData(String data) {
        super.writeData(encode(data));
    }

    @Override
    public String readData() {
        return decode(super.readData());
    }
}



























//  65070501001 Kantapong Vongpanich
//
//           .ooo     oooooooo   .oooo.    ooooooooo   .oooo.     oooooooo   .oooo.     .o    .oooo.     .oooo.     .o
//         .88'      dP"""""""  d8P'`Y8b  d"""""""8'  d8P'`Y8b   dP"""""""  d8P'`Y8b  o888   d8P'`Y8b   d8P'`Y8b  o888
//        d88'      d88888b.   888    888       .8'  888    888 d88888b.   888    888  888  888    888 888    888  888
//       d888P"Ybo.     `Y88b  888    888      .8'   888    888     `Y88b  888    888  888  888    888 888    888  888
//       Y88[   ]88       ]88  888    888     .8'    888    888       ]88  888    888  888  888    888 888    888  888
//       `Y88   88P o.   .88P  `88b  d88'    .8'     `88b  d88' o.   .88P  `88b  d88'  888  `88b  d88' `88b  d88'  888
//        `88bod8'  `8bd88P'    `Y8bd8P'    .8'       `Y8bd8P'  `8bd88P'    `Y8bd8P'  o888o  `Y8bd8P'   `Y8bd8P'  o888o
//
//
//
//       oooo    oooo                           .
//       `888   .8P'                          .o8
//        888  d8'     .oooo.   ooo. .oo.   .o888oo  .oooo.   oo.ooooo.   .ooooo.  ooo. .oo.    .oooooooo
//        88888[      `P  )88b  `888P"Y88b    888   `P  )88b   888' `88b d88' `88b `888P"Y88b  888' `88b
//        888`88b.     .oP"888   888   888    888    .oP"888   888   888 888   888  888   888  888   888
//        888  `88b.  d8(  888   888   888    888 . d8(  888   888   888 888   888  888   888  `88bod8P'
//       o888o  o888o `Y888""8o o888o o888o   "888" `Y888""8o  888bod8P' `Y8bod8P' o888o o888o `8oooooo.
//                                                             888                             d"     YD
//                                                            o888o                            "Y88888P'
//
//       oooooo     oooo                                                                    o8o            oooo
//        `888.     .8'                                                                     `"'            `888
//         `888.   .8'    .ooooo.  ooo. .oo.    .oooooooo oo.ooooo.   .oooo.   ooo. .oo.   oooo   .ooooo.   888 .oo.
//          `888. .8'    d88' `88b `888P"Y88b  888' `88b   888' `88b `P  )88b  `888P"Y88b  `888  d88' `"Y8  888P"Y88b
//           `888.8'     888   888  888   888  888   888   888   888  .oP"888   888   888   888  888        888   888
//            `888'      888   888  888   888  `88bod8P'   888   888 d8(  888   888   888   888  888   .o8  888   888
//             `8'       `Y8bod8P' o888o o888o `8oooooo.   888bod8P' `Y888""8o o888o o888o o888o `Y8bod8P' o888o o888o
//                                             d"     YD   888
//                                             "Y88888P'  o888o
//