load("@aspect_bazel_lib//lib:expand_template.bzl", "expand_template_rule")
load("@aspect_bazel_lib//lib:write_source_files.bzl", "write_source_files")

def generate_black(names):
    files = {}

    for name in names:
        template_name = name.replace("white", "template")
        expand_template_rule(
            name = template_name,
            out = template_name + ".tmpl",
            substitutions = {
                "White": "{Color}",
                "Black": "{OponentColor}",
            },
            template = name + ".go",
        )

        black_name = name.replace("white", "black")

        result_name = black_name + "_go"
        expand_template_rule(
            name = result_name,
            out = result_name + "gen",
            substitutions = {
                "{Color}": "Black",
                "{OponentColor}": "White",
            },
            template = template_name,
        )

        files[black_name + ".go"] = result_name

    source_files_name = "copy_black_source_files"
    write_source_files(
        name = source_files_name,
        files = files,
    )
