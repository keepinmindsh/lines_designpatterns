package lines.loader;

public class ClassLoader {

    public static void main(String[] args) {
        JVMComponent.Builder builder = new JVMComponent.Builder();


        builder.Heap(new Heap());
        builder.JavaThreads(new JavaThreads());
        builder.MethodArea(new MethodArea());
        builder.ProgramCounterRegisters(new ProgramCounterRegisters());
        builder.NativeInternalThreads(new NativeInternalThreads());

        ClassLoadSubSystem classLoadSubSystem = new ClassLoadSubSystem(builder.build());

        classLoadSubSystem.ControlVariables();
    }
}

class ClassLoadSubSystem {

    JVMComponent jvmComponent;

    public ClassLoadSubSystem(JVMComponent jvmComponent){
        this.jvmComponent = jvmComponent;
    }

    public void ControlVariables(){

        JVMElement methodArea = jvmComponent.getJVMCompoentOrNull(ComponentType.MethodArea);
        JVMElement heap = jvmComponent.getJVMCompoentOrNull(ComponentType.Heap);
        JVMElement javaThreads = jvmComponent.getJVMCompoentOrNull(ComponentType.JavaThreads);
        JVMElement nativeInternalThreads = jvmComponent.getJVMCompoentOrNull(ComponentType.NativeInternalThreads);
        JVMElement programCounterRegisters = jvmComponent.getJVMCompoentOrNull(ComponentType.ProgramCounterRegisters);


        methodArea.doProcess();
        heap.doProcess();
        javaThreads.doProcess();
        nativeInternalThreads.doProcess();
        programCounterRegisters.doProcess();
    }
}

abstract class JVMElement {
    public abstract void doProcess();
}

class Heap extends JVMElement {

    public void doProcess() {

        forHeapSetting();
    }

    private void forHeapSetting(){

    }
}

class JavaThreads extends JVMElement {
    @Override
    public void doProcess() {

        forJavaThreadSetting();
    }

    private void forJavaThreadSetting(){

    }
}

class MethodArea extends JVMElement {

    public void doProcess() {

        forMethodAreaSetting();
    }

    private void forMethodAreaSetting(){

    }
}

class NativeInternalThreads extends JVMElement {

    public void doProcess() {

        forNativeInternalThreadSetting();
    }

    private void forNativeInternalThreadSetting(){

    }
}

class ProgramCounterRegisters extends JVMElement {

    public void doProcess() {

        forProgramCounterRegisters();
    }

    private void forProgramCounterRegisters(){

    }
}

class JVMComponent {

    private final Heap heap;
    private final JavaThreads javaThreads;
    private final MethodArea methodArea;
    private final NativeInternalThreads nativeInternalThreads;
    private final ProgramCounterRegisters programCounterRegisters;

    static class Builder{

        private Heap heap;
        private JavaThreads javaThreads;
        private MethodArea methodArea;
        private NativeInternalThreads nativeInternalThreads;
        private ProgramCounterRegisters programCounterRegisters;

        public Builder Heap(Heap heap){
            this.heap = heap;
            return this;
        }

        public Builder JavaThreads(JavaThreads javaThreads){
            this.javaThreads = javaThreads;
            return this;
        }

        public Builder MethodArea(MethodArea methodArea){
            this.methodArea = methodArea;
            return this;
        }

        public Builder NativeInternalThreads(NativeInternalThreads nativeInternalThreads){
            this.nativeInternalThreads = nativeInternalThreads;
            return this;
        }

        public Builder ProgramCounterRegisters(ProgramCounterRegisters programCounterRegisters){
            this.programCounterRegisters = programCounterRegisters;
            return this;

        }


        public JVMComponent build(){
            return new JVMComponent(this);
        }

    }

    private JVMComponent(Builder build){
        this.heap = build.heap;
        this.javaThreads = build.javaThreads;
        this.methodArea = build.methodArea;
        this.nativeInternalThreads = build.nativeInternalThreads;
        this.programCounterRegisters = build.programCounterRegisters;
    }

    public JVMElement getJVMCompoentOrNull(ComponentType type){
        switch (type){
            case Heap :
                return heap;
            case JavaThreads:
                return javaThreads;
            case MethodArea:
                return methodArea;
            case NativeInternalThreads:
                return nativeInternalThreads;
            case ProgramCounterRegisters:
                return programCounterRegisters;
        }

        return null;
    }
}

enum ComponentType {
    Heap,
    JavaThreads,
    MethodArea,
    NativeInternalThreads,
    ProgramCounterRegisters
}