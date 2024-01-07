load("@aspect_bazel_lib//lib:expand_template.bzl", "expand_template_rule")
load("@aspect_bazel_lib//lib:write_source_files.bzl", "write_source_files")

def generate_black(name):
    template_name = name + "_template"
    expand_template_rule(
        name = template_name,
        out = template_name + ".tmpl",
        substitutions = {
            "White": "{Color}",
            "Black": "{OponentColor}",
        },
        template = name + "_white.go",
    )

    result_name = name + "_go"
    expand_template_rule(
        name = result_name,
        out = result_name + "gen",
        substitutions = {
            "{Color}": "Black",
            "{OponentColor}": "White",
        },
        template = template_name,
    )

    source_files_name = name + "_source_files"
    write_source_files(
        name = source_files_name,
        files = {
            name + "_black.go": result_name,
        },
        tags = ["workspace_copy"]
    )
