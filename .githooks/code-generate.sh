QUERY_PARAMETERS="--noshow_progress --ui_event_filters=-info,-stderr"
RUN_PARAMETERS="--remote_download_toplevel --noshow_progress --ui_event_filters=-info,-stdout,-stderr"

TEMPLATE_COPY_TARGETS=$(bazel query $QUERY_PARAMETERS 'attr(name, "^copy_black_source_files$", //...)' | grep . || true)
if [ -n "$TEMPLATE_COPY_TARGETS" ]; then
    echo generate bazel workspace files ...
    echo "$TEMPLATE_COPY_TARGETS" | xargs -L 1 bazel run $RUN_PARAMETERS
fi
