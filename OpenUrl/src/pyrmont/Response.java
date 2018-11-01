package pyrmont;

import java.io.File;  
import java.io.FileInputStream;  
import java.io.IOException;  
import java.io.OutputStream;  
  
/** 
 * HTTP Response = Status-Line 
 *      *(( general-header | response-header | entity-header ) CRLF)  
 *      CRLF 
 *      [message-body] 
 *      Status-Line=Http-Version SP Status-Code SP Reason-Phrase CRLF 
 * 
 */  
public class Response {  
    private static final int BUFFER_SIZE=1024;  
    Request request;  
    OutputStream output;  
      
    public Response(OutputStream output){  
        this.output=output;  
    }  
      
    public void setRequest(Request request){  
        this.request=request;  
    }  
      
    public void sendStaticResource()throws IOException{
String Message="HTTP/1.1 200 OK\r\n"+  
                "Content-Type:text/html\r\n"+  
                "Content-Length:53\r\n"+  
                "\r\n"+  
                "<h1>This is CICD test page!</h1>\r\n";  
                output.write(Message.getBytes());
        }  
}