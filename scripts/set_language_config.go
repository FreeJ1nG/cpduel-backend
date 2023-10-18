package scripts

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	AllowUse  bool
	Extension string
}

var languageConfig = map[string]string{
	"7":  "python",
	"31": "python",
	"32": "go",
	"40": "python",
	"41": "python",
	"50": "cpp",
	"54": "cpp",
	"60": "java",
	"70": "python",
	"73": "cpp",
	"74": "java",
	"87": "java",
}

func SetLanguageConfig(ctx context.Context, mainDB *pgxpool.Pool) {
	for id, extension := range languageConfig {
		fmt.Printf("Attempting to update language with id %s\n", id)
		_, err := mainDB.Exec(
			ctx,
			`UPDATE Language
			SET 
				allow_use = true,
				extension = $1
			WHERE
				id = $2`,
			extension,
			id,
		)
		if err != nil {
			fmt.Printf("Unable to update language with id %s: %s\n", id, err.Error())
			return
		}
		fmt.Printf("Finished updating language %s with extension of %s\n", id, extension)
	}
	fmt.Println("Script Finished!")
	os.Exit(0)
}

// 43 | GNU GCC C11 5.1.0                  | f         |
// 80 | Clang++20 Diagnostics              | f         |
// 52 | Clang++17 Diagnostics              | f         |
// 50 | GNU G++14 6.4.0                    | f         |
// 54 | GNU G++17 7.3.0                    | f         |
// 73 | GNU G++20 11.2.0 (64 bit, winlibs) | f         |
// 59 | Microsoft Visual C++ 2017          | f         |
// 61 | GNU G++17 9.2.0 (64 bit, msys 2)   | f         |
// 65 | C# 8, .NET Core 3.1                | f         |
// 79 | C# 10, .NET SDK 6.0                | f         |
// 9  | C# Mono 6.8                        | f         |
// 28 | D DMD32 v2.105.0                   | f         |
// 32 | Go 1.19.5                          | f         |
// 12 | Haskell GHC 8.10.1                 | f         |
// 60 | Java 11.0.6                        | f         |
// 74 | Java 17 64bit                      | f         |
// 87 | Java 21 64bit                      | f         |
// 36 | Java 1.8.0_241                     | f         |
// 77 | Kotlin 1.6.10                      | f         |
// 83 | Kotlin 1.7.20                      | f         |
// 19 | OCaml 4.02.1                       | f         |
// 3  | Delphi 7                           | f         |
// 4  | Free Pascal 3.0.2                  | f         |
// 51 | PascalABC.NET 3.8.3                | f         |
// 13 | Perl 5.20.1                        | f         |
// 6  | PHP 8.1.7                          | f         |
// 7  | Python 2.7.18                      | f         |
// 31 | Python 3.8.10                      | f         |
// 40 | PyPy 2.7.13 (7.3.0)                | f         |
// 41 | PyPy 3.6.9 (7.3.0)                 | f         |
// 70 | PyPy 3.9.10 (7.3.9, 64bit)         | f         |
// 67 | Ruby 3.2.2                         | f         |
// 75 | Rust 1.72.0 (2021)                 | f         |
// 20 | Scala 2.12.8                       | f         |
// 34 | JavaScript V8 4.8.0                | f         |
// 55 | Node.js 12.16.3                    | f         |
// 56 | Microsoft Q#                       | f         |
// 68 | Secret 2021                        | f         |
