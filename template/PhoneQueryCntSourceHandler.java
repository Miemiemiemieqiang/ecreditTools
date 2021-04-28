package com.ecreditpal.vodka.variable.handler.source;

import com.ecreditpal.vodka.variable.feign.entity.response.advance.AdvanceMultiPlatformComplexDetectionResq;
import com.ecreditpal.vodka.variable.handler.Invoke;
import com.ecreditpal.vodka.variable.handler.Source;
import com.ecreditpal.vodka.variable.handler.SourceHandler;


/**
 * Created by {{.Author}} on {{.Datetime}}
 */
@Source(value = "{{.Lower}}", alias = "{{.Lower}}")
public class {{.Upper}}SourceHandler extends SourceHandler {

    @Invoke
    public Integer getSource(AdvanceMultiPlatformComplexDetectionResq multiPlatformComplexDetection) {
        if (multiPlatformComplexDetection == null) {
            return null;
        }
        return multiPlatformComplexDetection.getPhoneStatistics();
    }
}
