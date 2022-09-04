micro run .
micro cap SayHello --message=wang
micro call cap Cap.SayHello '{"message": "wang"}'

micro run --name wff .
micro wff Cap SayHello --message=wang
micro call wff Cap.SayHello '{"message": "wang"}'