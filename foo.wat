(module
    (func $add (param $one i32) (param $two i32) (result i32)
        local.get $one
        local.get $two
        i32.add
    )
    (export "add" (func $add))
)