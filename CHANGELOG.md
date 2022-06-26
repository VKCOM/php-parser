# Changelog

All notable changes to this project will be documented in this file, in reverse chronological order by release.

## `v0.8.2` 2022-26-06

### Added

- [#28](https://github.com/VKCOM/php-parser/pull/28): `php8.2`: added readonly classes support
- [#29](https://github.com/VKCOM/php-parser/pull/29): `php8.1`: added intersection types support

## `v0.8.1-rc.1` 2021-09-08

### Added

- [#6](https://github.com/VKCOM/php-parser/pull/6): `php8.1`: added `readonly` modifier
- [#8](https://github.com/VKCOM/php-parser/pull/8): `php8.1`: added `never` type
- [#10](https://github.com/VKCOM/php-parser/pull/10): `php8.1`: added new octal numbers syntax
- [#12](https://github.com/VKCOM/php-parser/pull/12): `php8.1`: added enums
- [#15](https://github.com/VKCOM/php-parser/pull/15): `php8.1`: added `final` modifier for constants in class
- [#18](https://github.com/VKCOM/php-parser/pull/18): `php8.1`: added first class callable syntax

### Changed

- [`4cd50d`](https://github.com/VKCOM/php-parser/commit/85b5d3ef36c9b12923404caf1c57497aa84cd50d): `cmd`: added file path output before errors

### Fixed

- [#22](https://github.com/VKCOM/php-parser/pull/22): fixed bug with `#` comments

## `v0.8.0-rc.2` 2021-30-07

### Added

- [#10](https://github.com/i582/php-parser/pull/10): `php8`: nullsafe operator (`?->`)
- [#13](https://github.com/i582/php-parser/pull/13): `php8`: named arguments 
- [#19](https://github.com/i582/php-parser/pull/19): `php8`: `match` expression 
- [#21](https://github.com/i582/php-parser/pull/21): `php8`: union types in type hints and `static` typehint 
- [#23](https://github.com/i582/php-parser/pull/23): `php8`: block `catch` without variable 
- [#25](https://github.com/i582/php-parser/pull/25): `php8`: trailing comma in parameter lists 
- [#27](https://github.com/i582/php-parser/pull/27): `php8`: `throw` can be used as an expression 
- [#32](https://github.com/i582/php-parser/pull/32): `php8`: declaring properties in the constructor 
- [#34](https://github.com/i582/php-parser/pull/34): `php8`: attributes 
- [#38](https://github.com/i582/php-parser/pull/38): `php8`: trailing comma in closure use list 

### Changed

- [#30](https://github.com/i582/php-parser/pull/30): `php8`: concatenation precedence 
- [#36](https://github.com/i582/php-parser/pull/36): `php8`: names in the namespace are treated as a single token 
- [#42](https://github.com/i582/php-parser/pull/42): `php8`: deferencable changes and arbitrary expressions in `new`/`instanceof` 

### Removed

- [#11](https://github.com/i582/php-parser/pull/11): `php8`: removed `(real)` cast 
- [#15](https://github.com/i582/php-parser/pull/15): `php8`: removed `(unset)` cast 
- [#17](https://github.com/i582/php-parser/pull/17): `php8`: removed `{}` access 

---

Versions prior to 0.8.0 were not included in this changelog.