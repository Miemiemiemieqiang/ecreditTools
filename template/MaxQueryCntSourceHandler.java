package com.ecreditpal.vodka.variable.handler.source;

import com.ecreditpal.vodka.variable.feign.entity.response.advance.AdvanceMultiPlatformDetectionResp;
import com.ecreditpal.vodka.variable.handler.Invoke;
import com.ecreditpal.vodka.variable.handler.Source;
import com.ecreditpal.vodka.variable.handler.SourceHandler;

import static com.ecreditpal.vodka.variable.handler.source.rulefeature.advance.MaxQueryCnt6h2wkSourceHandler.getMaxQuery;


/**
 * Created by {{.Author}} on {{.Datetime}}
 */
@Source(value = "{{.Lower}}", alias = "{{.Lower}}")
public class {{.Upper}}SourceHandler extends SourceHandler {

    private static final String TIME_SLICE = "6-HOUR";

    @Invoke
    public Integer getSource(AdvanceMultiPlatformDetectionResp multiPlatformDetection) {
        return getMaxQuery(multiPlatformDetection, TIME_SLICE);
    }
}
