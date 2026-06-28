load("@npm//@bazel/typescript:index.bzl", "ts_library")
# Placeholder for testing and styling macros
# load("@bazel_concatjs//:index.bzl", "karma_web_test")
# load("@io_bazel_rules_sass//:defs.bzl", "sass_binary")

def ng_component(name, srcs, deps = [], has_spec = False, **kwargs):
    """
    A Bazel macro to reduce boilerplate when defining Angular components.
    It automatically defines a ts_library for the component and optionally
    sets up testing targets if a .spec.ts file exists.
    """
    ts_library(
        name = name,
        srcs = srcs,
        deps = deps,
        **kwargs
    )
    
    if has_spec:
        # Define the spec library
        spec_srcs = [s for s in srcs if s.endswith(".spec.ts")]
        if spec_srcs:
            ts_library(
                name = name + "_spec",
                srcs = spec_srcs,
                deps = deps + [":" + name],
                testonly = True,
            )
            # Future: karma_web_test(name = name + "_test", deps = [":" + name + "_spec"])
