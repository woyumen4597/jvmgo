public class MyObject{
	public static int staticVar;
	public int instanceVar;
	public static void main(String[] args) {
		int x = 327;
		MyObject myobj = new MyObject();
		MyObject.staticVar = x;
		x = MyObject.staticVar;
		myobj.instanceVar = x;
		x = myobj.instanceVar;
		Object obj = myobj;
		if(obj instanceof MyObject){
			myobj = (MyObject) obj;
			System.out.println(myobj.instanceVar);
		}
	}
}