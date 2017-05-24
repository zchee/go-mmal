// +build ignore

int
main(){
    int value;
    value = 0;
    
    int* ptr;
    ptr = &value;
    
    struct X{};
    X x;
    x;
    
    auto px = &x;
    px;
    return 0;
}
