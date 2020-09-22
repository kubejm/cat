open Unix

let buffer_size = 8192

let buffer = Bytes.create buffer_size

let cat filename =
  let fd = openfile filename [ O_RDONLY ] 0 in
  let rec write_stdout () =
    match read fd buffer 0 buffer_size with
    | 0 -> ()
    | n ->
        ignore (write stdout buffer 0 n : file_perm);
        write_stdout ()
  in
  write_stdout ();
  close fd

let () =
  for i = 1 to Array.length Sys.argv - 1 do
    cat Sys.argv.(i)
  done
